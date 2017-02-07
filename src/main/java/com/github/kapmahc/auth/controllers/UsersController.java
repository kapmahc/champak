package com.github.kapmahc.auth.controllers;

import com.auth0.jwt.JWT;
import com.auth0.jwt.interfaces.DecodedJWT;
import com.github.kapmahc.auth.errors.ForbiddenException;
import com.github.kapmahc.auth.forms.EmailForm;
import com.github.kapmahc.auth.forms.ResetPasswordForm;
import com.github.kapmahc.auth.forms.SignInForm;
import com.github.kapmahc.auth.forms.SignUpForm;
import com.github.kapmahc.auth.models.Log;
import com.github.kapmahc.auth.models.User;
import com.github.kapmahc.auth.repositories.UserRepository;
import com.github.kapmahc.auth.services.UserService;
import com.github.kapmahc.auth.utils.JwtHelper;
import com.github.kapmahc.auth.utils.SecurityUtil;
import org.msgpack.core.MessageBufferPacker;
import org.msgpack.core.MessagePack;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.context.MessageSource;
import org.springframework.security.access.annotation.Secured;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.stereotype.Controller;
import org.springframework.validation.BindingResult;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.servlet.mvc.support.RedirectAttributes;

import javax.annotation.Resource;
import javax.validation.Valid;
import java.io.IOException;
import java.security.GeneralSecurityException;
import java.util.Date;
import java.util.Locale;

/**
 * Created by flamen on 17-1-23.
 */
@Controller("auth.usersC")
@RequestMapping("/users")
public class UsersController {

    @GetMapping("/sign-in")
    public String getSignIn(SignInForm signInForm) {
        return "auth/users/sign-in";
    }

    @PostMapping("/sign-in")
    public String postSignIn(@Valid SignInForm form, BindingResult bindingResult) {
        //todo
        if (bindingResult.hasErrors()) {
            return "auth/users/sign-in";
        }
        return "redirect:/";
    }

    @GetMapping("/sign-up")
    public String getSignUp(SignUpForm form) {
        return "auth/users/sign-up";
    }

    @PostMapping("/sign-up")
    public String postSignUp(@Valid SignUpForm form, BindingResult bindingResult, RedirectAttributes redirectAttributes, Locale locale) throws IOException, GeneralSecurityException {
        if (!bindingResult.hasErrors()) {
            if (form.getPassword().equals(form.getPasswordConfirmation())) {
                if (userRepository.findByEmail(form.getEmail()) == null) {
                    User user = userService.add(form.getName(), form.getEmail(), form.getPassword());
                    userService.log(user, messageSource.getMessage("auth.logs.sign-up", null, locale));
                    sendEmail(user, Action.CONFIRM, locale);
                    redirectAttributes.addFlashAttribute("notice", messageSource.getMessage("auth.messages.receive-confirm-email", null, locale));
                    return "redirect:/users/sign-in";
                } else {
                    bindingResult.rejectValue("email", "auth.messages.email-already-exist");
                }
            } else {
                bindingResult.rejectValue("password", "messages.passwords-not-match");
            }
        }
        return "auth/users/sign-up";
    }

    @GetMapping("/unlock/{token}")
    public String getUnlock(@PathVariable String token, RedirectAttributes redirectAttributes, Locale locale) {
        User user = parseToken(token, Action.UNLOCK);
        if (user == null) {
            redirectAttributes.addFlashAttribute("alert", messageSource.getMessage("auth.messages.user-not-exist", null, locale));
        } else {
            if (user.getLockedAt() == null) {
                redirectAttributes.addFlashAttribute("alert", messageSource.getMessage("auth.messages.user-not-lock", null, locale));
            } else {
                user.setLockedAt(null);
                userRepository.save(user);
                userService.log(user, messageSource.getMessage("auth.log.unlock", null, locale));
                redirectAttributes.addFlashAttribute("notice", messageSource.getMessage("auth.messages.unlock-success", null, locale));
            }
        }
        return "redirect:/users/sign-in";
    }

    @GetMapping("/unlock")
    public String getUnlock(EmailForm form) {
        return "auth/users/unlock";
    }

    @PostMapping("/unlock")
    public String postUnlock(@Valid EmailForm form, BindingResult bindingResult, RedirectAttributes redirectAttributes, Locale locale) throws IOException, GeneralSecurityException {
        if (!bindingResult.hasErrors()) {
            User user = userRepository.findByEmail(form.getEmail());
            if (user == null) {
                bindingResult.rejectValue("email", "auth.messages.user-not-exist");
            } else {
                if (user.getLockedAt() == null) {
                    bindingResult.rejectValue("email", "auth.messages.user-not-lock");
                } else {
                    sendEmail(user, Action.UNLOCK, locale);
                    redirectAttributes.addFlashAttribute("notice", messageSource.getMessage("auth.messages.receive-unlock-email", null, locale));
                    return "redirect:/users/sign-in";
                }
            }
        }
        return "auth/users/unlock";
    }

    @GetMapping("/confirm/{token}")
    public String getConfirm(@PathVariable String token, RedirectAttributes redirectAttributes, Locale locale) {
        User user = parseToken(token, Action.CONFIRM);
        if (user == null) {
            redirectAttributes.addFlashAttribute("alert", messageSource.getMessage("auth.messages.user-not-exist", null, locale));
        } else {
            if (user.getConfirmedAt() == null) {
                user.setConfirmedAt(new Date());
                userRepository.save(user);
                userService.log(user, messageSource.getMessage("auth.log.confirm", null, locale));
                redirectAttributes.addFlashAttribute("notice", messageSource.getMessage("auth.messages.confirm-success", null, locale));
            } else {
                redirectAttributes.addFlashAttribute("alert", messageSource.getMessage("auth.messages.user-already-confirm", null, locale));
            }
        }
        return "redirect:/users/sign-in";
    }

    @GetMapping("/confirm")
    public String getConfirm(EmailForm form) {
        return "auth/users/confirm";
    }

    @PostMapping("/confirm")
    public String postConfirm(@Valid EmailForm form, BindingResult bindingResult, RedirectAttributes redirectAttributes, Locale locale) throws IOException, GeneralSecurityException {
        if (!bindingResult.hasErrors()) {
            User user = userRepository.findByEmail(form.getEmail());
            if (user == null) {
                bindingResult.rejectValue("email", "auth.messages.user-not-exist");
            } else {
                if (user.getConfirmedAt() == null) {
                    sendEmail(user, Action.CONFIRM, locale);
                    redirectAttributes.addFlashAttribute("notice", messageSource.getMessage("auth.messages.receive-confirm-email", null, locale));
                    return "redirect:/users/sign-in";
                } else {
                    bindingResult.rejectValue("email", "auth.messages.user-already-confirm");
                }
            }
        }
        return "auth/users/confirm";
    }

    @GetMapping("/forgot-password")
    public String getForgotPassword(EmailForm form) {
        return "auth/users/forgot-password";
    }

    @PostMapping("/forgot-password")
    public String postForgotPassword(@Valid EmailForm form, BindingResult bindingResult, RedirectAttributes redirectAttributes, Locale locale) throws IOException, GeneralSecurityException {
        if (!bindingResult.hasErrors()) {
            User user = userRepository.findByEmail(form.getEmail());
            if (user == null) {
                bindingResult.rejectValue("email", "auth.messages.user-not-exist");
            } else {
                if (user.getConfirmedAt() == null) {
                    sendEmail(user, Action.RESET_PASSWORD, locale);
                    redirectAttributes.addFlashAttribute("notice", messageSource.getMessage("auth.messages.receive-reset-password-email", null, locale));
                    return "redirect:/users/sign-in";
                } else {
                    bindingResult.rejectValue("email", "auth.messages.user-already-confirm");
                }
            }
        }
        return "auth/users/forgot-password";
    }

    @GetMapping("/reset-password/{token}")
    public String getChangePassword(ResetPasswordForm form, @PathVariable String token) {
        form.setToken(token);
        return "auth/users/reset-password";
    }

    @PostMapping("/reset-password")
    public String postChangePassword(@Valid ResetPasswordForm form, BindingResult bindingResult, RedirectAttributes redirectAttributes, Locale locale) {
        if (!bindingResult.hasErrors()) {
            if (form.getPassword().equals(form.getPasswordConfirmation())) {
                User user = parseToken(form.getToken(), Action.RESET_PASSWORD);
                if (user == null) {
                    bindingResult.rejectValue("email", "auth.messages.user-not-exist");
                } else {
                    user.setPassword(securityUtil.password(form.getPassword()));
                    userRepository.save(user);
                    userService.log(user, messageSource.getMessage("auth.logs.reset-password", null, locale));
                    redirectAttributes.addFlashAttribute("notice", messageSource.getMessage("auth.messages.reset-password-success", null, locale));
                    return "redirect:/users/sign-in";
                }
            } else {
                bindingResult.rejectValue("password", "messages.passwords-not-match");
            }
        }
        return "auth/users/reset-password";
    }

    enum Action {
        CONFIRM, UNLOCK, RESET_PASSWORD
    }

    private User parseToken(String token, Action action) {
        DecodedJWT jwt = jwtHelper.parse(token);
        if (!"auth".equals(jwt.getClaim("mod").asString()) || action != Action.valueOf(jwt.getClaim("act").asString())) {
            throw new ForbiddenException();
        }
        return userRepository.findByEmail(jwt.getClaim("uid").asString());
    }

    private void sendEmail(User user, Action action, Locale locale) throws IOException, GeneralSecurityException {
        String token = jwtHelper.generate(
                JWT.create().withClaim("uid", user.getUid()).withClaim("act", action.name()).withClaim("mod", "auth"),
                1);
        String path;
        switch (action) {
            case CONFIRM:
                path = "confirm";
                break;
            case UNLOCK:
                path = "unlock";
                break;
            case RESET_PASSWORD:
                path = "reset-password";
                break;
            default:
                throw new ForbiddenException();
        }

        MessageBufferPacker packer = MessagePack.newDefaultBufferPacker();
        packer
                .packString(user.getEmail())
                .packString(messageSource.getMessage(
                        String.format("auth.emails.%s.subject", path),
                        null,
                        locale))
                .packString(messageSource.getMessage(
                        String.format("auth.emails.%s.body", path),
                        new Object[]{
                                user.getFullName(),
                                String.format(
                                        "https://%s/users/%s/%s",
                                        settingService.get("site.domain",String.class),
                                        path,
                                        token
                                ),
                        },
                        locale));
        packer.close();

        jobSender.send("email", packer.toByteArray());
    }

    @Resource
    MessageSource messageSource;
    @Resource
    UserRepository userRepository;
    @Resource
    UserService userService;
    @Resource
    SecurityUtil securityUtil;
    @Resource
    JwtHelper jwtHelper;

    private static final Logger logger = LoggerFactory.getLogger(UsersController.class);
}
