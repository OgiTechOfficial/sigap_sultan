import {
  AUTO_STYLE,
  animate,
  state,
  style,
  transition,
  trigger,
} from '@angular/animations';

export const openCloseProfileMenuAnimation = trigger('openCloseProfile', [
  state(
    'true',
    style({
      opacity: 1,
      transform: 'scale(1)',
    })
  ),
  state(
    'false',
    style({
      opacity: 0,
      transform: 'scale(.95)',
    })
  ),
  transition('false => true', animate('100ms ease-out')),
]);

export const rotateAnimation = trigger('rotateMenu', [
  state('true', style({ transform: 'rotate(180deg)' })),
  transition('false <=> true', animate(`250ms ease-out`)),
]);

export const openCloseMobileMenu = trigger('openCloseMobileMenu', [
  state('true', style({ height: AUTO_STYLE, display: 'flex' })),
  state('false', style({ height: 0, display: 'none' })),
  transition('false <=> true', animate(`250ms ease-in`))
])
