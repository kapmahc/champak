package com.github.kapmahc.auth.init;

import com.github.kapmahc.auth.utils.SignOutHandler;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.config.annotation.authentication.builders.AuthenticationManagerBuilder;
import org.springframework.security.config.annotation.method.configuration.EnableGlobalMethodSecurity;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter;
import org.springframework.security.config.http.SessionCreationPolicy;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.session.data.redis.config.annotation.web.http.EnableRedisHttpSession;
import org.springframework.web.servlet.config.annotation.EnableWebMvc;

import javax.annotation.Resource;

/**
 * Created by flamen on 17-1-24.
 */
@Configuration
@EnableRedisHttpSession
@EnableWebSecurity
public class SecurityConfig extends WebSecurityConfigurerAdapter {
    @Override
    protected void configure(AuthenticationManagerBuilder auth) throws Exception {
        auth.userDetailsService(userDetailsService);
    }

    @Override
    protected void configure(HttpSecurity http) throws Exception {
        http
                .sessionManagement().sessionCreationPolicy(SessionCreationPolicy.STATELESS)
                .and().authorizeRequests().antMatchers("/druid/**", "/monitoring").hasRole("admin")
                .and().formLogin().loginPage("/users/sign-in").defaultSuccessUrl("/dashboard")
                .and().logout().logoutUrl("/users/sign-out").logoutSuccessUrl("/").addLogoutHandler(signOutHandler).invalidateHttpSession(true)
                .and().csrf().ignoringAntMatchers("/druid/**", "/monitoring")
        ;
    }

    @Resource
    UserDetailsService userDetailsService;
    @Resource
    SignOutHandler signOutHandler;
}
