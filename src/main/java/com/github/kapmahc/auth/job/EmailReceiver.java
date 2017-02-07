package com.github.kapmahc.auth.job;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Component;

/**
 * Created by flamen on 17-1-24.
 */
@Component("auth.emailReceiver")
public class EmailReceiver {
    public void receiveMessage(String message){
        logger.debug("Received: {}", message);
    }
    private static final Logger logger = LoggerFactory.getLogger(EmailReceiver.class);
}
