package angular_bazel_grpc.multiplication_server;

import multiplication_service.MultiplicationServiceGrpc;
import multiplication_service.ProductRequest;
import multiplication_service.ProductResponse;
import addition_service.AdditionServiceGrpc;
import addition_service.SumRequest;
import addition_service.SumResponse;
import io.grpc.ManagedChannelBuilder;
import io.grpc.stub.StreamObserver;

public class MultiplicationServiceImpl extends MultiplicationServiceGrpc.MultiplicationServiceImplBase {
  private final AdditionServiceGrpc.AdditionServiceBlockingStub additionServiceBlockingStub =
      AdditionServiceGrpc.newBlockingStub(ManagedChannelBuilder
                                              .forAddress("localhost", 9090)
                                              // Channels are secure by default (via SSL/TLS). For the example we
                                              // disable TLS to avoid needing certificates.
                                              .usePlaintext()
                                              .build());

  @Override
  public void product(ProductRequest req, StreamObserver<ProductResponse> responseObserver) {
    SumResponse sumResponse =
        additionServiceBlockingStub.sum(SumRequest.newBuilder().addAllIntegers(req.getIntegersList()).build());
    ProductResponse response = ProductResponse.newBuilder().setProduct(sumResponse.getSum()).build();
    responseObserver.onNext(response);
    responseObserver.onCompleted();
  }
}