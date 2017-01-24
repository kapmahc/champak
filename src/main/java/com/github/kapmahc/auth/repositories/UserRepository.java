package com.github.kapmahc.auth.repositories;

import com.github.kapmahc.auth.models.User;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

/**
 * Created by flamen on 17-1-23.
 */
@Repository("auth.userR")
public interface UserRepository extends CrudRepository<User, Long> {
}
