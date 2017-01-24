package com.github.kapmahc.auth.repositories;

import com.github.kapmahc.auth.models.Log;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

/**
 * Created by flamen on 17-1-24.
 */
@Repository("auth.logR")
public interface LogRepository extends CrudRepository<Log, Long> {
}
