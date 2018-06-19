import {ChangeDetectionStrategy, Component} from '@angular/core';
import {AdditionServiceService, AdditionServiceSumRequest} from 'angular_bazel_GRPC/addition_service/swagger_gen';
import {MultiplicationServiceProductRequest, MultiplicationServiceService} from 'angular_bazel_GRPC/multiplication_service/swagger_gen';

import {Observable} from 'rxjs';
import {map} from 'rxjs/operators';

const sumRequest: AdditionServiceSumRequest = {
  integers: [1, 2, 3],
};

const productRequest: MultiplicationServiceProductRequest = {
  integers: [2, 3, 4],
};

@Component({
  selector: 'app-component',
  templateUrl: 'app.component.html',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AppComponent {
  readonly sum$: Observable<number> =
      this.additionService.sum(sumRequest).pipe(map(response => response.sum));

  readonly product$: Observable<number> =
      this.multiplicationService.product(productRequest).pipe(map(response => response.product));

  constructor(
      private readonly additionService: AdditionServiceService,
      private readonly multiplicationService: MultiplicationServiceService,
  ) {}
}
