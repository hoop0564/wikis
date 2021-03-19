package frontcontroller.pattern;

public class FrontController {
    private Dispatcher dispatcher;

    public FrontController() {
        dispatcher = new Dispatcher();
    }

    private boolean isAuthenticUser() {
        System.out.println("user is authenticated successfully.");
        return true;
    }

    private void traceRequest(String request) {
        System.out.println("request = " + request);
    }

    public void dispatchRequest(String request) {
        traceRequest(request);
        if (isAuthenticUser()){
            dispatcher.dispatch(request);
        }
    }
}
