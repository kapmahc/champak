package com.github.kapmahc.auth.services;

import com.github.kapmahc.auth.models.User;
import com.github.kapmahc.auth.repositories.UserRepository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Component;

import javax.annotation.Resource;
import java.util.ArrayList;
import java.util.List;

/**
 * Created by flamen on 17-1-24.
 */
@Component("auth.userDetailsService")
public class UserDetailsService implements org.springframework.security.core.userdetails.UserDetailsService {
    @Override
    public UserDetails loadUserByUsername(String uid) throws UsernameNotFoundException {
        logger.debug("sign in {}", uid);

        User user = userRepository.findByEmail(uid);
        if (user == null) {
            throw new UsernameNotFoundException("not found");
        }
        List<SimpleGrantedAuthority> authorities = new ArrayList<>();
//        authorities.add(new SimpleGrantedAuthority(user.getRole().name()));

        return new org.springframework.security.core.userdetails.User(
                user.getUid(),
                user.getPassword(),
                authorities);
    }

    @Resource
    UserRepository userRepository;
    private final static Logger logger = LoggerFactory.getLogger(UserDetailsService.class);
}
