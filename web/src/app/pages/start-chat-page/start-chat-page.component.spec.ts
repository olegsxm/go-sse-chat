import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StartChatPageComponent } from './start-chat-page.component';

describe('StartChatPageComponent', () => {
  let component: StartChatPageComponent;
  let fixture: ComponentFixture<StartChatPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [StartChatPageComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(StartChatPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
