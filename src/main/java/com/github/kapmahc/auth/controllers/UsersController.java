package com.github.kapmahc.auth.controllers;

import com.github.kapmahc.auth.forms.SignInForm;
import com.github.kapmahc.auth.repositories.UserRepository;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;

import javax.annotation.Resource;

/**
 * Created by flamen on 17-1-23.
 */
@Controller("auth.usersC")
@RequestMapping("/users")
public class UsersController {

    @GetMapping("/sign-in")
    public String signIn(SignInForm signInForm) {
        return "auth/users/sign-in";
    }

    @PreAuthorize("permitAll()")
    @GetMapping("/sign-up")
    public String signUp() {
        return "auth/users/sign-up";
    }

    @Resource
    UserRepository userRepository;
}
