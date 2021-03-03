package datastruct;

import java.util.Enumeration;
import java.util.Vector;

public class EnumerationTester {
    public static void main(String[] args) {
        Enumeration<String> days;
        Vector<String> dayNames = new Vector<String>();
        dayNames.add("Monday");
        dayNames.add("Tuesday");
        dayNames.add("Wednesday");
        days = dayNames.elements();
        while (days.hasMoreElements()) {
            System.out.println(days.nextElement());
        }

        for (int i = 0; i < 5; i++) {
            
        }
    }
}