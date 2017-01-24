package com.github.kapmahc.auth.controllers;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;

/**
 * Created by flamen on 17-1-24.
 */
@Controller("auth.errorsC")
public class ErrorsController {
    @RequestMapping("/403")
    public String forbidden() {
        return "403";
    }
}
