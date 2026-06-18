part of '../daemon_client.dart';

String _payloadExample(String name) {
  switch (name) {
    case 'open_ott':
      return '{"service":"netflix"}';
    case 'open_url':
      return '{"url":"https://www.youtube.com"}';
    case 'open_youtube_search':
      return '{"query":"lofi music"}';
    case 'volume_set':
      return '{"level":30}';
    case 'volume_up':
    case 'volume_down':
      return '{"step":10}';
    default:
      return '{}';
  }
}

IconData _commandIcon(String name) {
  switch (name) {
    case 'open_coupang_play':
    case 'open_netflix':
      return Icons.movie_filter_outlined;
    case 'open_disney_plus':
      return Icons.auto_awesome_outlined;
    case 'open_tving':
      return Icons.live_tv_outlined;
    case 'open_wavve':
      return Icons.waves_outlined;
    case 'open_youtube':
    case 'open_youtube_search':
      return Icons.play_circle_outline;
    case 'open_ott':
    case 'movie_mode':
      return Icons.theaters_outlined;
    case 'volume_set':
    case 'volume_up':
    case 'volume_down':
      return Icons.volume_up_outlined;
    case 'volume_mute':
      return Icons.volume_off_outlined;
    case 'sleep_mode':
    case 'mac_sleep':
      return Icons.bedtime_outlined;
    case 'display_sleep':
      return Icons.monitor_outlined;
    default:
      return Icons.terminal_outlined;
  }
}
