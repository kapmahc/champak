package com.github.kapmahc.auth.repositories;

import com.github.kapmahc.auth.models.Locale;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

/**
 * Created by flamen on 17-1-24.
 */
@Repository("auth.localeR")
public interface LocaleRepository extends CrudRepository<Locale, Long> {
    Locale findByLangAndCode(String lang, String code);
}
