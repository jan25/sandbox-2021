package com.example;

/**
 * Adds two numbers.
 */
public class Add {
    public static <T extends Number> Number add(T a, T b) {
        return a.doubleValue() + b.doubleValue();
    }
}
