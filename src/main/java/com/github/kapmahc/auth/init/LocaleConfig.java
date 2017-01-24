package com.github.kapmahc.auth.init;

import com.github.kapmahc.auth.utils.ExposedResourceBundleMessageSource;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.LocaleResolver;
import org.springframework.web.servlet.config.annotation.*;
import org.springframework.web.servlet.i18n.CookieLocaleResolver;
import org.springframework.web.servlet.i18n.LocaleChangeInterceptor;
import org.springframework.web.servlet.resource.ContentVersionStrategy;
import org.springframework.web.servlet.resource.VersionResourceResolver;

import java.util.Locale;

/**
 * Created by flamen on 17-1-24.
 */
@Configuration
@EnableWebMvc
public class LocaleConfig extends WebMvcConfigurerAdapter {
    @Bean(name = "auth.filesMessageSource")
    public ExposedResourceBundleMessageSource getMessageSource() {
        //can't ust classpath:messages
        ExposedResourceBundleMessageSource ms = new ExposedResourceBundleMessageSource();
        ms.setBasename(msBasename);
        ms.setDefaultEncoding("UTF-8");
        ms.setCacheSeconds(msCacheSeconds);
        ms.setUseCodeAsDefaultMessage(true);
        ms.setFallbackToSystemLocale(true);
        return ms;
    }

    @Bean
    public LocaleResolver localeResolver() {
        CookieLocaleResolver clr = new CookieLocaleResolver();
        clr.setDefaultLocale(Locale.US);
        clr.setCookieName(LOCALE);
        clr.setCookieMaxAge(Integer.MAX_VALUE);
        clr.setLanguageTagCompliant(true);
        return clr;
    }

    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        LocaleChangeInterceptor interceptor = new LocaleChangeInterceptor();
        interceptor.setParamName(LOCALE);
        interceptor.setLanguageTagCompliant(true);
        registry.addInterceptor(interceptor);
    }

    @Override
    public void addResourceHandlers(ResourceHandlerRegistry registry) {
        VersionResourceResolver vrr = new VersionResourceResolver().addVersionStrategy(new ContentVersionStrategy(), "/**");

        registry.addResourceHandler("/statics/**")
                .addResourceLocations("classpath:/themes/" + theme + "/statics/")
                .setCachePeriod(60 * 60 * 24 * 365) /* one year */
                .resourceChain(true)
                .addResolver(vrr);
    }

    @Override
    public void configureDefaultServletHandling(DefaultServletHandlerConfigurer configurer) {
        configurer.enable();
    }

    private final static String LOCALE = "locale";
    @Value("${app.theme}")
    protected String theme;
    @Value("${spring.messages.basename}")
    protected String msBasename;
    @Value("${spring.messages.cache-seconds}")
    protected int msCacheSeconds;
}
