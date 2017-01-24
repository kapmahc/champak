package com.github.kapmahc.auth.utils;

import org.springframework.context.support.ResourceBundleMessageSource;

import java.util.Locale;
import java.util.ResourceBundle;
import java.util.Set;

/**
 * Created by flamen on 17-1-24.
 */
public class ExposedResourceBundleMessageSource extends ResourceBundleMessageSource {
    public Set<String> getKeys(String basename, Locale locale) {
        ResourceBundle bundle = getResourceBundle(basename, locale);
        return bundle.keySet();
    }
}
