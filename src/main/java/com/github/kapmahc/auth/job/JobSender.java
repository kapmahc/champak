package com.github.kapmahc.auth.job;

import com.mysql.cj.mysqlx.io.MessageBuilder;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.data.redis.connection.Message;
import org.springframework.stereotype.Component;

import javax.annotation.Resource;
import java.util.UUID;

/**
 * Created by flamen on 17-1-24.
 */
@Component("auth.jobSender")
public class JobSender {
    public void send(String act, byte[] body) {
        send(act, MessageProperties.CONTENT_TYPE_BYTES, body);
    }

    public void send(String act, String body) {
        send(act, MessageProperties.CONTENT_TYPE_TEXT_PLAIN, body.getBytes());
    }

    private void send(String act, String type, byte[] body) {
        Message msg = MessageBuilder
                .withBody(body)
                .setContentType(type)
                .setMessageId(UUID.randomUUID().toString())
                .setHeader("act", act).build();
        template.send(queueName, msg);
    }

    @Value("#{'${app.name}'+'.jobs'}")
    String queueName;
    @Resource
    AmqpTemplate template;

    public String getQueueName() {
        return queueName;
    }
}
