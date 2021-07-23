package com.example;

public class App {
    public static void main(String[] args) {
        if (args.length < 2) {
            System.err.println("Enter two numbers to add, e.g. 3 2.4");
            return;
        }
        System.out.println(
                "Result: " + String.valueOf(Add.add(Double.parseDouble(args[0]), Double.parseDouble(args[1]))));
    }
}
