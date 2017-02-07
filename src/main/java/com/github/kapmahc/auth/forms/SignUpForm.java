package com.github.kapmahc.auth.forms;

import org.hibernate.validator.constraints.Email;
import org.hibernate.validator.constraints.NotEmpty;

import javax.validation.constraints.Size;
import java.io.Serializable;

/**
 * Created by flamen on 17-1-24.
 */
public class SignUpForm implements Serializable {
    @Email
    @NotEmpty
    @Size(max = 255)
    private String email;
    @NotEmpty
    @Size(max = 255)
    private String name;
    @NotEmpty
    @Size(min = 6, max = 255)
    private String password;
    private String passwordConfirmation;

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
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
