package com.github.kapmahc;

import com.github.kapmahc.auth.models.User;
import com.github.kapmahc.auth.services.UserService;
import com.github.kapmahc.auth.utils.SecurityUtil;
import org.junit.Assert;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import javax.annotation.Resource;
import java.util.UUID;

@RunWith(SpringRunner.class)
@SpringBootTest
public class ApplicationTests {

    @Test
    public void contextLoads() {
    }

    @Test
    public void testDao() {
        for (int i = 0; i < 5; i++) {
            User user = userService.add("test", String.format("%s@test.com", UUID.randomUUID().toString()), "change-me");
            Assert.assertNotNull(user.getId());
        }
    }

    @Test
    public void testSecurity() {
        String hello = "Hello. CHAMPAK!";
        for (int i = 0; i < 5; i++) {
            String password = securityUtil.password(hello);
            System.out.printf("password('%s') = %s\n", hello, password);
            Assert.assertTrue(securityUtil.check(hello, password));

            String code = securityUtil.encrypt(hello);
            System.out.printf("encrypt('%s') = %s\n", hello, code);
            Assert.assertEquals(hello, securityUtil.decrypt(code));

        }
    }

    @Resource
    SecurityUtil securityUtil;
    @Resource
    UserService userService;

}
