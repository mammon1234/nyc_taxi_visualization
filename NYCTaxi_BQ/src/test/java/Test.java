import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;

/**
 * Created by YangYu on 4/18/16.
 */
public class Test {

    private static SimpleDateFormat sfg = new SimpleDateFormat("MM/dd/yyyy hh:mm a");

    public static void main(String[] args) {
        Date date = null;
        try {
            date = sfg.parse("04/18/2016 12:00 AM");
        } catch (ParseException e) {
            e.printStackTrace();
        }
        System.out.println(date);
    }

}
