package servlet;

import bigquery.BigQueryClient;
import com.google.api.services.bigquery.model.TableCell;
import com.google.api.services.bigquery.model.TableRow;
import com.google.gson.Gson;
import com.google.gson.GsonBuilder;

import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@WebServlet(name = "FareServlet")
public class FareServlet extends HttpServlet {

    protected void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        doGet(request, response);
    }

    protected void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
//        String sql = request.getParameter("sql");
        String sql = "SELECT\n" +
                "  total_amount\n" +
                "FROM\n" +
                "  [NYC.Trips]\n" +
                "WHERE\n" +
                "  /* Return values between a pair of */ /* latitude and longitude coordinates */ pickup_latitude > \"40.755932\"\n" +
                "  AND pickup_latitude < \"40.761783\"\n" +
                "  AND pickup_longitude > \"-73.983615\"\n" +
                "  AND pickup_longitude < \"-73.987735\"\n" +
                "  AND dropoff_latitude > \"40.639577\"\n" +
                "  AND dropoff_latitude < \"40.659737\"\n" +
                "  AND dropoff_longitude > \"-73.774269\"\n" +
                "  AND dropoff_longitude < \"-73.791908\"\n";
        List<TableRow> rows = BigQueryClient.executeQuery(sql);
        List<Double> fare = new ArrayList<Double>();
        Map<String, Object> map = new HashMap<String, Object>();

        if (rows != null) {
            map.put("result", "success");
            for (TableRow row : rows) {
                for (TableCell field : row.getF()) {
                    String v = (String) field.getV();
                    fare.add(Double.valueOf(v));
                }
            }
        } else {
            map.put("result", "fail");
        }
        map.put("list", fare);
        Gson gson = new GsonBuilder().create();
        String json = gson.toJson(map);
        response.setContentType("application/json");
        response.setCharacterEncoding("UTF-8");
        response.getWriter().write(json);
    }
}
