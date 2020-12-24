import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MasterNavbarComponent } from './master-navbar.component';

describe('MasterNavbarComponent', () => {
  let component: MasterNavbarComponent;
  let fixture: ComponentFixture<MasterNavbarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MasterNavbarComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(MasterNavbarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
