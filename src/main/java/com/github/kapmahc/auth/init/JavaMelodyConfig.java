package com.github.kapmahc.auth.init;

import net.bull.javamelody.MonitoringFilter;
import net.bull.javamelody.Parameter;
import net.bull.javamelody.SessionListener;
import org.springframework.boot.web.servlet.FilterRegistrationBean;
import org.springframework.boot.web.servlet.ServletContextInitializer;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.ImportResource;

import javax.servlet.DispatcherType;
import javax.servlet.ServletContext;
import javax.servlet.ServletException;

/**
 * Created by flamen on 17-1-23.
 */
@Configuration
@ImportResource("classpath:net/bull/javamelody/monitoring-spring.xml")
public class JavaMelodyConfig implements ServletContextInitializer {
    @Override
    public void onStartup(ServletContext servletContext) throws ServletException {
        servletContext.addListener(new SessionListener());
    }

    @Bean
    public FilterRegistrationBean javaMelody() {
        final FilterRegistrationBean javaMelody = new FilterRegistrationBean();
        final MonitoringFilter filter = new MonitoringFilter();
        filter.setApplicationType("Spring Boot");
        javaMelody.setFilter(filter);
        javaMelody.setAsyncSupported(true);
        javaMelody.setName("javamelody");
        javaMelody.setDispatcherTypes(DispatcherType.REQUEST, DispatcherType.ASYNC);

        // https://github.com/javamelody/javamelody/wiki/UserGuide#6-optional-parameters
        javaMelody.addInitParameter(Parameter.LOG.getCode(), Boolean.toString(true));

        javaMelody.addUrlPatterns("/*");
        return javaMelody;
    }
}
