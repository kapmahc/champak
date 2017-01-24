package com.github.kapmahc.auth.utils;

import com.github.kapmahc.auth.repositories.LocaleRepository;
import org.springframework.context.support.AbstractMessageSource;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import javax.annotation.Resource;
import java.text.MessageFormat;
import java.util.Locale;

/**
 * Created by flamen on 17-1-24.
 */
@Component("messageSource")
public class DatabaseMessageSource extends AbstractMessageSource {

    @Override
    protected MessageFormat resolveCode(String code, Locale locale) {
        com.github.kapmahc.auth.models.Locale l = localeRepository.findByLangAndCode(locale.toLanguageTag(), code);
        return l == null ? null : createMessageFormat(l.getMessage(), locale);
    }

    @PostConstruct
    void init() {
        setParentMessageSource(messageSource);
    }

    @Resource
    LocaleRepository localeRepository;
    @Resource(name = "auth.filesMessageSource")
    ExposedResourceBundleMessageSource messageSource;

}
