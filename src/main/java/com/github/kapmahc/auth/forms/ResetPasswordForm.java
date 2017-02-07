package com.github.kapmahc.auth.forms;

import org.hibernate.validator.constraints.NotEmpty;

import javax.validation.constraints.Size;
import java.io.Serializable;

/**
 * Created by flamen on 17-1-24.
 */
public class ResetPasswordForm implements Serializable {
    @NotEmpty
    private String token;
    @NotEmpty
    @Size(min = 6, max=255)
    private String password;
    private String passwordConfirmation;

    public String getToken() {
        return token;
    }

    public void setToken(String token) {
        this.token = token;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public String getPasswordConfirmation() {
        return passwordConfirmation;
    }

    public void setPasswordConfirmation(String passwordConfirmation) {
        this.passwordConfirmation = passwordConfirmation;
    }
}
