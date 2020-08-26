import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PhaseTenComponent } from './phase-ten.component';

describe('PhaseTenComponent', () => {
  let component: PhaseTenComponent;
  let fixture: ComponentFixture<PhaseTenComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PhaseTenComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PhaseTenComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
