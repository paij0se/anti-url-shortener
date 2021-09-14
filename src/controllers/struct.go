package controllers

type proxyList struct {
	Data []proxyData `json:"data"`
}
type proxyData struct {
	ID        string   `json:"_id"`
	IP        string   `json:"ip"`
	Port      string   `json:"port"`
	Protocols []string `json:"protocols"`
}

/*
{
  _id: '61409661840f815e069936ac',
  ip: '200.32.80.106',
  anonymityLevel: 'transparent',
  asn: 'AS3549',
  city: 'Bogotá',
  country: 'CO',
  created_at: '2021-09-14T12:32:33.199Z',
  google: false,
  isp: 'Level 3 Colombia S.A',
  lastChecked: 1631624639,
  latency: 117,
  org: 'UB@NET',
  port: '8080',
  protocols: [ 'http' ],
  region: null,
  responseTime: 35,
  speed: null,
  updated_at: '2021-09-14T13:03:59.279Z',
  workingPercent: null,
  upTime: 100,
  upTimeSuccessCount: 1,
  upTimeTryCount: 1
}
*/
