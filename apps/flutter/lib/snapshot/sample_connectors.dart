part of '../snapshot.dart';

const _sampleConnectors = [
  ConnectorReadiness(
    key: 'mydata',
    label: 'MyData aggregator',
    category: 'finance_aggregation',
    status: 'planned',
    fixtureMode: true,
    dataClasses: ['accounts', 'cards', 'transactions'],
    allowedOperations: ['read_fixture', 'summarize'],
    forbiddenOperations: [
      'credential_request',
      'external_api_call',
      'transfer',
      'trade',
      'card_action',
    ],
    nextStep:
        'Define consent and local vault boundaries before any real connector.',
  ),
  ConnectorReadiness(
    key: 'commerce',
    label: 'Commerce purchases',
    category: 'commerce',
    status: 'planned',
    fixtureMode: true,
    dataClasses: ['orders', 'items', 'recurring_candidates'],
    allowedOperations: ['read_fixture', 'recommend_review', 'summarize'],
    forbiddenOperations: [
      'credential_request',
      'cookie_import',
      'scraping',
      'purchase',
    ],
    nextStep:
        'Extend local purchase fixtures and avoid scraping/cookie capture.',
  ),
];
