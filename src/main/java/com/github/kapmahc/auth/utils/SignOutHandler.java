package com.github.kapmahc.auth.utils;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.security.core.Authentication;
import org.springframework.security.web.authentication.logout.LogoutHandler;
import org.springframework.stereotype.Component;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

/**
 * Created by flamen on 17-1-24.
 */
@Component("auth.signOutH")
public class SignOutHandler implements LogoutHandler {
    @Override
    public void logout(HttpServletRequest request, HttpServletResponse response, Authentication auth) {
        logger.info("{} sign out", auth.getName());
    }

    private final static Logger logger = LoggerFactory.getLogger(SignOutHandler.class);
}
