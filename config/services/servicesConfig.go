// Copyright 2018 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package services

import (
	"github.com/gohugoio/hugo/config"
	"github.com/mitchellh/mapstructure"
)

const (
	servicesConfigKey = "services"

	disqusShortnameKey = "disqusshortname"
	googleAnalyticsKey = "googleanalytics"
)

// Config is a privacy configuration for all the relevant services in Hugo.
type Config struct {
	Disqus          Disqus
	GoogleAnalytics GoogleAnalytics
}

// Disqus holds the functional configuration settings related to the Disqus template.
type Disqus struct {
	// A Shortname is the unique identifier assigned to a Disqus site.
	Shortname string
}

// GoogleAnalytics holds the functional configuration settings related to the Google Analytics template.
type GoogleAnalytics struct {
	// The GA tracking ID.
	ID string
}

func DecodeConfig(cfg config.Provider) (c Config, err error) {
	m := cfg.GetStringMap(servicesConfigKey)

	err = mapstructure.WeakDecode(m, &c)

	// Keep backwards compability.
	if c.GoogleAnalytics.ID == "" {
		// Try the global config
		c.GoogleAnalytics.ID = cfg.GetString(googleAnalyticsKey)
	}
	if c.Disqus.Shortname == "" {
		c.Disqus.Shortname = cfg.GetString(disqusShortnameKey)
	}

	return
}
