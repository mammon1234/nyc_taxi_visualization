package bigquery;

import com.google.api.client.googleapis.auth.oauth2.GoogleCredential;
import com.google.api.client.http.HttpTransport;
import com.google.api.client.http.javanet.NetHttpTransport;
import com.google.api.client.json.JsonFactory;
import com.google.api.client.json.jackson2.JacksonFactory;
import com.google.api.services.bigquery.Bigquery;
import com.google.api.services.bigquery.BigqueryScopes;
import com.google.api.services.bigquery.model.GetQueryResultsResponse;
import com.google.api.services.bigquery.model.QueryRequest;
import com.google.api.services.bigquery.model.QueryResponse;
import com.google.api.services.bigquery.model.TableCell;
import com.google.api.services.bigquery.model.TableRow;

import java.io.IOException;
import java.util.List;



public class BigQueryClient {

    private final static String projectId = "fast-pagoda-127101";

    /**
     * Creates an authorized Bigquery client service using Application Default Credentials.
     *
     * @return an authorized Bigquery client
     * @throws IOException if there's an error getting the default credentials.
     */
    private static Bigquery createAuthorizedClient() throws IOException {
        // Create the credential
        HttpTransport transport = new NetHttpTransport();
        JsonFactory jsonFactory = new JacksonFactory();
        GoogleCredential credential = GoogleCredential.getApplicationDefault(transport, jsonFactory);

        if (credential.createScopedRequired()) {
            credential = credential.createScoped(BigqueryScopes.all());
        }

        return new Bigquery.Builder(transport, jsonFactory, credential)
                .setApplicationName("Bigquery Samples").build();
    }

    public static List<TableRow> executeQuery(String querySql) {
        try {
            Bigquery bigquery = createAuthorizedClient();
            QueryResponse query = bigquery.jobs().query(
                    projectId,
                    new QueryRequest().setQuery(querySql))
                    .execute();
//            query.set("allowLargeResults", true);
            // Execute it
            GetQueryResultsResponse queryResult = bigquery.jobs().getQueryResults(
                    query.getJobReference().getProjectId(),
                    query.getJobReference().getJobId()).execute();
            return queryResult.getRows();
        } catch (IOException e) {
            return null;
        }
    }

    /**
     * Prints the results to standard out.
     *
     * @param rows the rows to print.
     */
    private static void printResults(List<TableRow> rows) {
        System.out.print("\nQuery Results:\n------------\n");
        for (TableRow row : rows) {
            for (TableCell field : row.getF()) {
                System.out.printf("%-50s", field.getV());
            }
            System.out.println();
        }
    }
}
