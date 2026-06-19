part of '../snapshot.dart';

const _sampleOpenCommands = [
  HomeCommand(
    name: 'open-youtube-search',
    payload: '{"query":"lofi music"}',
    icon: Icons.play_circle_outline,
    payloadFields: ['query'],
  ),
  HomeCommand(
    name: 'open-url',
    payload: '{"url":"https://www.youtube.com"}',
    icon: Icons.public_outlined,
    payloadFields: ['url'],
  ),
  HomeCommand(
    name: 'open-ott',
    payload: '{"service":"netflix"}',
    icon: Icons.theaters_outlined,
    payloadFields: ['service'],
  ),
];
