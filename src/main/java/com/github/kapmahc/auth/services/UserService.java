package com.github.kapmahc.auth.services;

import com.github.kapmahc.auth.models.Log;
import com.github.kapmahc.auth.models.User;
import com.github.kapmahc.auth.repositories.LogRepository;
import com.github.kapmahc.auth.repositories.UserRepository;
import com.github.kapmahc.auth.utils.SecurityUtil;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.UUID;

/**
 * Created by flamen on 17-1-24.
 */
@Service("auth.userS")
public class UserService {

    public void log(User user, String message) {
        Log l = new Log();
        l.setUser(user);
        l.setLevel(Log.Level.INFO);
        l.setMessage(message);
        logRepository.save(l);
    }

    public void log(User user, Log.Level level, String message) {
        Log l = new Log();
        l.setUser(user);
        l.setLevel(level);
        l.setMessage(message);
        logRepository.save(l);
    }

    public User add(String name, String email, String password) {
        User u = new User();
        u.setName(name);
        u.setEmail(email);
        u.setPassword(securityUtil.password(password));
        u.setUid(UUID.randomUUID().toString());
        u.setProviderId(email);
        u.setProviderType(User.Type.EMAIL);
        userRepository.save(u);
        return u;
    }

    @Resource
    UserRepository userRepository;
    @Resource
    LogRepository logRepository;
    @Resource
    SecurityUtil securityUtil;
}
