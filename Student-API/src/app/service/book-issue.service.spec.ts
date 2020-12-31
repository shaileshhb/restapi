import { TestBed } from '@angular/core/testing';

import { BookIssueService } from './book-issue.service';

describe('BookIssueService', () => {
  let service: BookIssueService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(BookIssueService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
