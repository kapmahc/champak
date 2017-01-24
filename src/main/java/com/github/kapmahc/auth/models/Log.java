package com.github.kapmahc.auth.models;

import org.hibernate.annotations.CreationTimestamp;
import org.hibernate.annotations.GenericGenerator;

import javax.persistence.*;
import java.io.Serializable;
import java.util.Date;

/**
 * Created by flamen on 17-1-23.
 */
@Entity
@Table(name = "logs", indexes = {
        @Index(columnList = "level")
})
public class Log implements Serializable {
    public enum Level {
        INFO, ERROR
    }

    @Id
    @GenericGenerator(
            name = "logsSequenceGenerator",
            strategy = "org.hibernate.id.enhanced.SequenceStyleGenerator"
    )
    @GeneratedValue(generator = "logsSequenceGenerator")
    private long id;
    @Column(nullable = false, updatable = false)
    private String message;
    @Column(nullable = false, updatable = false, length = 8)
    @Enumerated(EnumType.STRING)
    private Level level;
    @Temporal(TemporalType.TIMESTAMP)
    @CreationTimestamp
    @Column(nullable = false)
    private Date createdAt;
    @JoinColumn(nullable = false, updatable = false)
    @ManyToOne
    private User user;

    public long getId() {
        return id;
    }

    public void setId(long id) {
        this.id = id;
    }

    public User getUser() {
        return user;
    }

    public void setUser(User user) {
        this.user = user;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    public Level getLevel() {
        return level;
    }

    public void setLevel(Level level) {
        this.level = level;
    }

    public Date getCreatedAt() {
        return createdAt;
    }

    public void setCreatedAt(Date createdAt) {
        this.createdAt = createdAt;
    }
}
