package com.github.kapmahc.auth.controllers;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;

/**
 * Created by flamen on 17-1-24.
 */
@Controller("auth.homeC")
public class HomeController {
    @GetMapping("/dashboard")
    public String dashboard() {
        return "dashboard";
    }
}
