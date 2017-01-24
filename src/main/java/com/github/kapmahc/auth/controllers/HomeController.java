package com.github.kapmahc.auth.controllers;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;

/**
 * Created by flamen on 17-1-24.
 */
@Controller("auth.homeC")
public class HomeController {
    @GetMapping("/dashboard")
    public String dashboard() {
        return "dashboard";
    }

    @GetMapping("/")
    public String home() {
        return "home";
    }

    @PostMapping("/search")
    public String search() {
        return "search";
    }
}
