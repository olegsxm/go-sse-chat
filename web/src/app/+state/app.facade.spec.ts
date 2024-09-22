import { NgModule } from '@angular/core';
import { TestBed } from '@angular/core/testing';
import { EffectsModule } from '@ngrx/effects';
import { StoreModule, Store } from '@ngrx/store';
import { readFirst } from '@nx/angular/testing';

import * as AppActions from './app.actions';
import { AppEffects } from './app.effects';
import { AppFacade } from './app.facade';
import { AppEntity } from './app.models';
import {
  APP_FEATURE_KEY,
  AppState,
  initialAppState,
  appReducer,
} from './app.reducer';
import * as AppSelectors from './app.selectors';

interface TestSchema {
  app: AppState;
}

describe('AppFacade', () => {
  let facade: AppFacade;
  let store: Store<TestSchema>;
  const createAppEntity = (id: string, name = ''): AppEntity => ({
    id,
    name: name || `name-${id}`,
  });

  describe('used in NgModule', () => {
    beforeEach(() => {
      @NgModule({
        imports: [
          StoreModule.forFeature(APP_FEATURE_KEY, appReducer),
          EffectsModule.forFeature([AppEffects]),
        ],
        providers: [AppFacade],
      })
      class CustomFeatureModule {}

      @NgModule({
        imports: [
          StoreModule.forRoot({}),
          EffectsModule.forRoot([]),
          CustomFeatureModule,
        ],
      })
      class RootModule {}
      TestBed.configureTestingModule({ imports: [RootModule] });

      store = TestBed.inject(Store);
      facade = TestBed.inject(AppFacade);
    });

    /**
     * The initially generated facade::loadAll() returns empty array
     */
    it('loadAll() should return empty list with loaded == true', async () => {
      let list = await readFirst(facade.allApp$);
      let isLoaded = await readFirst(facade.loaded$);

      expect(list.length).toBe(0);
      expect(isLoaded).toBe(false);

      facade.init();

      list = await readFirst(facade.allApp$);
      isLoaded = await readFirst(facade.loaded$);

      expect(list.length).toBe(0);
      expect(isLoaded).toBe(true);
    });

    /**
     * Use `loadAppSuccess` to manually update list
     */
    it('allApp$ should return the loaded list; and loaded flag == true', async () => {
      let list = await readFirst(facade.allApp$);
      let isLoaded = await readFirst(facade.loaded$);

      expect(list.length).toBe(0);
      expect(isLoaded).toBe(false);

      store.dispatch(
        AppActions.loadAppSuccess({
          app: [createAppEntity('AAA'), createAppEntity('BBB')],
        })
      );

      list = await readFirst(facade.allApp$);
      isLoaded = await readFirst(facade.loaded$);

      expect(list.length).toBe(2);
      expect(isLoaded).toBe(true);
    });
  });
});
