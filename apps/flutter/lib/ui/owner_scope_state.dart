part of '../main.dart';

enum OwnerScopeState { ownerScoped, householdScoped, empty }

extension OwnerScopeStateLabel on OwnerScopeState {
  String get label => switch (this) {
    OwnerScopeState.ownerScoped => 'owner scoped',
    OwnerScopeState.householdScoped => 'household scoped',
    OwnerScopeState.empty => 'empty',
  };

  JarvisBadgeTone get tone => switch (this) {
    OwnerScopeState.ownerScoped => JarvisBadgeTone.secondary,
    OwnerScopeState.householdScoped => JarvisBadgeTone.secondary,
    OwnerScopeState.empty => JarvisBadgeTone.outline,
  };
}

List<OwnerScopeState> ownerScopeStates(String owner, int records) => [
  if (owner == 'household')
    OwnerScopeState.householdScoped
  else
    OwnerScopeState.ownerScoped,
  if (records == 0) OwnerScopeState.empty,
];

Color ownerScopeColor(String owner, int records) {
  if (records == 0) {
    return JarvisAstryxTokens.warning;
  }
  return owner == 'household'
      ? JarvisAstryxTokens.success
      : JarvisAstryxTokens.accent;
}
