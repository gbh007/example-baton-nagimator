local config = {
  dashboard: {
    title: 'Example baton nagimator',
    uid: '',
    timeFrom: 'now-1h',
    refresh: '1m',
  },
  variable: {
    prometheus: {
      type: 'prometheus',
      name: 'metrics',
    },
    host: {
      name: 'host',
    },
  },
  label: {
    filter: {
      service: 'instance=~"$%s", service_name="%s"' % [
        $.variable.host.name,
        'example-button-nagimator',
      ],
    },
    servicePrefix: 'buttoners_example_baton_nagimator_',
  },
  datasource: {
    metrics: {
      type: $.variable.prometheus.type,
      uid: '${%s}' % $.variable.prometheus.name,
    },
  },
};

local grafonnet = import 'github.com/grafana/grafonnet/gen/grafonnet-latest/main.libsonnet';

local variable = grafonnet.dashboard.variable;
local panel = grafonnet.panel;
local query = grafonnet.query;
local prometheus = grafonnet.query.prometheus;


local simpleTSLegend() =
  panel.timeSeries.options.legend.withDisplayMode('table')
  + panel.timeSeries.options.legend.withPlacement('bottom')
  + panel.timeSeries.options.legend.withCalcs(['mean', 'lastNotNull'])
  + panel.timeSeries.options.legend.withSortBy('Mean')
  + panel.timeSeries.options.legend.withSortDesc();

grafonnet.dashboard.new(config.dashboard.title)
+ grafonnet.dashboard.withUid(config.dashboard.uid)
+ grafonnet.dashboard.time.withFrom(config.dashboard.timeFrom)
+ grafonnet.dashboard.withRefresh(config.dashboard.refresh)
+ grafonnet.dashboard.graphTooltip.withSharedCrosshair()
+ grafonnet.dashboard.withPanels(
  grafonnet.util.grid.makeGrid([
    panel.timeSeries.new('Кнопочный РПС')
    + panel.timeSeries.queryOptions.withTargets([
      prometheus.new(
        config.datasource.metrics.uid,
        'sum(rate(%sbutton_lovers_total{%s}[$__rate_interval]))' % [
          config.label.servicePrefix,
          config.label.filter.service,
        ],
      ),
    ])
    + panel.timeSeries.options.legend.withShowLegend(false)
    + panel.timeSeries.standardOptions.withUnit('reqps')
    + panel.timeSeries.queryOptions.withDatasource(
      config.datasource.metrics.type,
      config.datasource.metrics.uid,
    ),
    panel.pieChart.new('Главные часовые нажиматоры')
    + panel.pieChart.queryOptions.withTargets([
      prometheus.new(
        config.datasource.metrics.uid,
        'sum(delta(%sbutton_lovers_total{%s}[1h])) by (name)' % [
          config.label.servicePrefix,
          config.label.filter.service,
        ],
      )
      + prometheus.withLegendFormat('{{name}}')
      + prometheus.withInstant(),
    ])
    + panel.pieChart.options.legend.withDisplayMode('table')
    + panel.pieChart.options.legend.withPlacement('right')
    + panel.pieChart.options.legend.withValues('value')
    + panel.pieChart.standardOptions.withUnit('short')
    + panel.pieChart.queryOptions.withDatasource(
      config.datasource.metrics.type,
      config.datasource.metrics.uid,
    ),
    panel.timeSeries.new('Я кушаю')
    + panel.timeSeries.queryOptions.withTargets([
      prometheus.new(
        config.datasource.metrics.uid,
        'sum(go_memstats_sys_bytes{%s})' % [
          config.label.filter.service,
        ],
      ),
    ])
    + panel.timeSeries.standardOptions.withUnit('bytes')
    + panel.timeSeries.options.legend.withShowLegend(false)
    + panel.timeSeries.queryOptions.withDatasource(
      config.datasource.metrics.type,
      config.datasource.metrics.uid,
    ),
  ], 8, 8, 0),
)
+ grafonnet.dashboard.withVariables([
  variable.datasource.new(
    config.variable.prometheus.name,
    config.variable.prometheus.type,
  ),
  variable.query.new(
    config.variable.host.name,
    'label_values(go_info{}, instance)',
  )
  + variable.query.withDatasource(config.datasource.metrics.type, config.datasource.metrics.uid)
  + variable.query.selectionOptions.withMulti()
  + variable.query.selectionOptions.withIncludeAll(true, '.*')
  + variable.query.refresh.onTime(),
])
