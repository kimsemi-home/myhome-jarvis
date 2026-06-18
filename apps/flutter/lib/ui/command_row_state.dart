part of '../main.dart';

class _CommandRowState extends State<CommandRow> {
  final _payload = _CommandPayloadState();

  @override
  void initState() {
    super.initState();
    _payload.sync(widget.command);
  }

  @override
  void didUpdateWidget(CommandRow oldWidget) {
    super.didUpdateWidget(oldWidget);
    if (_commandChanged(oldWidget.command, widget.command)) {
      _payload.sync(widget.command);
    }
  }

  @override
  void dispose() {
    _payload.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return CommandRowContent(
      command: widget.command,
      payload: _payload,
      onServiceChanged: (value) => setState(() => _payload.service = value),
      onDryRun: () => widget.onDryRun(_payload.edited(widget.command)),
    );
  }

  bool _commandChanged(HomeCommand oldCommand, HomeCommand nextCommand) {
    return oldCommand.name != nextCommand.name ||
        oldCommand.payload != nextCommand.payload ||
        oldCommand.payloadFields != nextCommand.payloadFields;
  }
}
