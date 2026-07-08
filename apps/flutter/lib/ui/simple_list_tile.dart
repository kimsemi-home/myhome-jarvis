part of '../main.dart';

class SimpleListTile extends StatelessWidget {
  const SimpleListTile({super.key, required this.title, required this.item});

  final String title;
  final String item;

  @override
  Widget build(BuildContext context) {
    final state = simpleListItemState(title, item);
    final iconColor = switch (state) {
      SimpleListItemState.offline => JarvisAstryxTokens.warning,
      SimpleListItemState.synced => JarvisAstryxTokens.success,
      SimpleListItemState.queued => JarvisAstryxTokens.warning,
      SimpleListItemState.configured => JarvisAstryxTokens.success,
      SimpleListItemState.localFixture => JarvisAstryxTokens.accent,
      SimpleListItemState.verified => JarvisAstryxTokens.success,
      SimpleListItemState.format => JarvisAstryxTokens.accent,
    };
    return Semantics(
      label: '$title item: $item, ${state.label}',
      child: JarvisSurface(
        padding: const EdgeInsets.all(14),
        child: Row(
          children: [
            Icon(Icons.circle_outlined, color: iconColor),
            const SizedBox(width: 12),
            Expanded(
              child: Text(item, maxLines: 2, overflow: TextOverflow.ellipsis),
            ),
            const SizedBox(width: 10),
            JarvisBadge(state.label, tone: state.tone),
          ],
        ),
      ),
    );
  }
}
