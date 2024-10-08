import {TestBed} from '@angular/core/testing';
import {provideStore, Store} from '@ngxs/store';
import {ChatState, IChatState} from './chat.state';
import {ChatAction} from './chat.actions';

describe('Chat store', () => {
    let store: Store;
    beforeEach(() => {
        TestBed.configureTestingModule({
            providers: [provideStore([ChatState])]

        });

        store = TestBed.inject(Store);
    });

    it('should create an action and add an item', () => {
        const expected: IChatState = {
            items: ['item-1']
        };
        store.dispatch(new ChatAction('item-1'));
        const actual = store.selectSnapshot(ChatState.getState);
        expect(actual).toEqual(expected);
    });

});
