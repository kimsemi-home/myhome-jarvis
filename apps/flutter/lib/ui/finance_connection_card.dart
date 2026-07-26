part of '../main.dart';

class FinanceConnectionCard extends StatelessWidget {
  const FinanceConnectionCard({
    super.key,
    required this.fixtureOnly,
    required this.dashboard,
  });

  final bool fixtureOnly;
  final FinanceDashboard dashboard;

  @override
  Widget build(BuildContext context) {
    final fixture = fixtureOnly;
    return JarvisSurface(
      padding: const EdgeInsets.all(14),
      child: Row(
        children: [
          const Icon(Icons.account_balance_outlined),
          const SizedBox(width: 12),
          Expanded(child: _connectionText(context, fixture)),
          JarvisBadge(
            fixture ? '샘플 데이터' : '연결됨',
            tone: fixture ? JarvisBadgeTone.muted : JarvisBadgeTone.success,
          ),
        ],
      ),
    );
  }

  Widget _connectionText(BuildContext context, bool fixture) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text('MyData 통합 조회', style: Theme.of(context).textTheme.titleSmall),
        const SizedBox(height: 3),
        Text(
          '토스처럼 계좌·카드·거래를 모아보는 연결 영역',
          style: Theme.of(context).textTheme.bodySmall,
        ),
        Text(
          fixture
              ? '로컬 fixture · 읽기 전용 · 정규화 패리티 ${dashboard.connectorParityPercent.toStringAsFixed(0)}%'
              : 'provider connected',
          style: Theme.of(context).textTheme.labelSmall,
        ),
      ],
    );
  }
}
