package angular_bazel_grpc.multiplication_server;

import addition_service.AdditionServiceGrpc;
import addition_service.SumRequest;
import io.grpc.ManagedChannelBuilder;
import io.grpc.stub.StreamObserver;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.IntStream;
import multiplication_service.MultiplicationServiceGrpc;
import multiplication_service.ProductRequest;
import multiplication_service.ProductResponse;

public class MultiplicationServiceImpl
    extends MultiplicationServiceGrpc.MultiplicationServiceImplBase {
  private final AdditionServiceGrpc.AdditionServiceBlockingStub additionServiceBlockingStub =
      AdditionServiceGrpc.newBlockingStub(
          ManagedChannelBuilder.forAddress("localhost", 9090)
              // Channels are secure by default (via SSL/TLS). For the example we
              // disable TLS to avoid needing certificates.
              .usePlaintext()
              .build());

  @Override
  public void product(ProductRequest req, StreamObserver<ProductResponse> responseObserver) {
    List<Integer> integers = req.getIntegersList();
    Integer multiplier = integers.get(0);
    for (int i = 0; i < integers.size() - 1; i++) {
      final Integer value = integers.get(i + 1);
      multiplier =
          additionServiceBlockingStub
              .sum(
                  SumRequest.newBuilder()
                      .addAllIntegers(
                          IntStream.generate(() -> value)
                              .limit(multiplier)
                              .boxed()
                              .collect(Collectors.<Integer>toList()))
                      .build())
              .getSum();
    }

    ProductResponse response = ProductResponse.newBuilder().setProduct(multiplier).build();
    responseObserver.onNext(response);
    responseObserver.onCompleted();
  }
}
