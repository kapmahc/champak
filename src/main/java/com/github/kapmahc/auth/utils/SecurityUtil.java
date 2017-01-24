package com.github.kapmahc.auth.utils;

import org.jasypt.util.password.PasswordEncryptor;
import org.jasypt.util.password.StrongPasswordEncryptor;
import org.jasypt.util.text.StrongTextEncryptor;
import org.jasypt.util.text.TextEncryptor;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;

/**
 * Created by flamen on 17-1-24.
 */
@Component("auth.securityUtil")
public class SecurityUtil {

    public String encrypt(String plain) {
        return te.encrypt(plain);
    }

    public String decrypt(String code) {
        return te.decrypt(code);
    }

    public String password(String plain) {
        return pe.encryptPassword(plain);
    }

    public boolean check(String plain, String code) {
        return pe.checkPassword(plain, code);
    }

    @PostConstruct
    void init() {
        pe = new StrongPasswordEncryptor();
        StrongTextEncryptor ste = new StrongTextEncryptor();
        ste.setPassword(password);
        te = ste;
    }

    @Value("${app.secrets.jce}")
    String password;

    private PasswordEncryptor pe;
    private TextEncryptor te;
}
