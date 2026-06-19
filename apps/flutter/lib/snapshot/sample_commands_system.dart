part of '../snapshot.dart';

const _sampleSystemCommands = [
  HomeCommand(
    name: 'volume-set',
    payload: '{"level":30}',
    icon: Icons.volume_up_outlined,
    payloadFields: ['level'],
  ),
  HomeCommand(
    name: 'volume-up',
    payload: '{"step":10}',
    icon: Icons.volume_up_outlined,
    payloadFields: ['step'],
  ),
  HomeCommand(
    name: 'volume-down',
    payload: '{"step":10}',
    icon: Icons.volume_down_outlined,
    payloadFields: ['step'],
  ),
  HomeCommand(
    name: 'volume-mute',
    payload: '{}',
    icon: Icons.volume_off_outlined,
  ),
  HomeCommand(
    name: 'display-sleep',
    payload: '{}',
    icon: Icons.monitor_outlined,
  ),
  HomeCommand(name: 'mac-sleep', payload: '{}', icon: Icons.bedtime_outlined),
  HomeCommand(name: 'movie-mode', payload: '{}', icon: Icons.theaters_outlined),
  HomeCommand(name: 'sleep-mode', payload: '{}', icon: Icons.bedtime_outlined),
];
