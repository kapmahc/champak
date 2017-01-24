package com.github.kapmahc.auth.controllers;

import com.github.kapmahc.auth.repositories.UserRepository;
import org.springframework.stereotype.Controller;

import javax.annotation.Resource;

/**
 * Created by flamen on 17-1-23.
 */
@Controller("auth.UsersController")
public class UsersController {
    @Resource
    UserRepository userRepository;
}
