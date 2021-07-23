package com.example;

import junit.framework.TestCase;

public class AddTest extends TestCase {
    public void testAdd() {
        assertTrue("add(2, 3) is 5.0", (double) Add.add(2, 3) == 5.0);
        assertTrue("add(3, 6.2) is 9.2", (double) Add.add(3, 6.2) == 9.2);
    }
}
