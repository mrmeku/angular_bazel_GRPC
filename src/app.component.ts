import {ChangeDetectionStrategy, Component} from '@angular/core';
import {AdditionServiceService, AdditionServiceSumRequest} from 'angular_bazel_GRPC/addition_service/swagger_gen';
import {Observable} from 'rxjs';
import {map} from 'rxjs/operators';

const request: AdditionServiceSumRequest = {
  integers: [1, 2, 3],
};

@Component({
  selector: 'app-component',
  templateUrl: 'app.component.html',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AppComponent {
  sum$: Observable<number> = this.additionService.sum(request).pipe(map(response => response.sum));

  constructor(private readonly additionService: AdditionServiceService) {}
}
