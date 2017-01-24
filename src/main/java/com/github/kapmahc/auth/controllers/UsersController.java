package com.github.kapmahc.auth.controllers;

import com.github.kapmahc.auth.repositories.UserRepository;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.web.authentication.logout.SecurityContextLogoutHandler;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;

import javax.annotation.Resource;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

/**
 * Created by flamen on 17-1-23.
 */
@Controller("auth.usersC")
@RequestMapping("/users")
public class UsersController {
    @GetMapping("/sign-in")
    public String signIn() {
        return "users/sign-in";
    }

    @DeleteMapping("/sign-out")
    public void signOut(HttpServletRequest request, HttpServletResponse response){
        Authentication auth = SecurityContextHolder.getContext().getAuthentication();
        if (auth != null){
            new SecurityContextLogoutHandler().logout(request, response, auth);
        }
    }

    @GetMapping("/sign-up")
    public String signUp() {
        return "users/sign-up";
    }

    @Resource
    UserRepository userRepository;
}
