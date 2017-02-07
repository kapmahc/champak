package com.github.kapmahc.auth.init;

import com.github.kapmahc.auth.job.EmailReceiver;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.redis.connection.RedisConnectionFactory;
import org.springframework.data.redis.listener.PatternTopic;
import org.springframework.data.redis.listener.RedisMessageListenerContainer;
import org.springframework.data.redis.listener.adapter.MessageListenerAdapter;

import javax.annotation.Resource;

/**
 * Created by flamen on 17-1-24.
 */
@Configuration
public class JobConfig {
    @Bean
    RedisMessageListenerContainer container(RedisConnectionFactory connectionFactory) {

        RedisMessageListenerContainer container = new RedisMessageListenerContainer();
        container.setConnectionFactory(connectionFactory);
        container.addMessageListener(new MessageListenerAdapter(emailReceiver, "receiveMessage"), new PatternTopic("emails"));

        return container;
    }

    @Resource
    EmailReceiver emailReceiver;
}
