part of '../main.dart';

class HouseholdView extends StatefulWidget {
  const HouseholdView({super.key, required this.scopes});

  final List<HouseholdScope> scopes;

  @override
  State<HouseholdView> createState() => _HouseholdViewState();
}

class _HouseholdViewState extends State<HouseholdView> {
  String? _selectedScope;

  @override
  void initState() {
    super.initState();
    _selectedScope = _firstScope(widget.scopes);
  }

  @override
  void didUpdateWidget(HouseholdView oldWidget) {
    super.didUpdateWidget(oldWidget);
    final scopes = widget.scopes.map((scope) => scope.scope).toSet();
    if (_selectedScope == null || !scopes.contains(_selectedScope)) {
      _selectedScope = _firstScope(widget.scopes);
    }
  }

  @override
  Widget build(BuildContext context) {
    if (widget.scopes.isEmpty) {
      return const PlainListView(title: 'Household', items: ['No scope data']);
    }
    final selected = widget.scopes.firstWhere(
      (scope) => scope.scope == _selectedScope,
      orElse: () => widget.scopes.first,
    );
    return HouseholdScopeBody(
      scopes: widget.scopes,
      selected: selected,
      onSelectionChanged: (selection) {
        setState(() {
          _selectedScope = selection.first;
        });
      },
    );
  }

  String? _firstScope(List<HouseholdScope> scopes) {
    return scopes.isEmpty ? null : scopes.first.scope;
  }
}
