package manage

import (
	"encoding/json"
	"testing"
)

func TestTypeTravisWebhookBody(t *testing.T) {
	testString := []byte(`
		{
		  "id": 289215159,
		  "number": "2848",
		  "config": {
		    "sudo": false,
		    "dist": "trusty",
		    "language": "python",
		    "python": [
		      "3.5.2"
		    ],
		    "branches": {
		      "only": [
		        "master"
		      ]
		    },
		    "cache": {
		      "pip": true,
		      "directories": [
		        "vendor/bundle",
		        "node_modules"
		      ]
		    },
		    "deploy": {
		      "provider": "heroku",
		      "api_key": {
		        "secure": "hylw2GIHMvZKOKX3uPSaLEzVrUGEA9mzGEA0s4zK37W9HJCTnvAcmgRCwOkRuC4L7R4Zshdh/CGORNnBBgh1xx5JGYwkdnqtjHuUQmWEXCusrIURu/iEBNSsZZEPK7zBuwqMHj2yRm64JfbTDJsku3xdoA5Z8XJG5AMJGKLFgUQ="
		      },
		      "app": "docs-travis-ci-com",
		      "skip_cleanup": true,
		      "true": {
		        "branch": [
		          "master"
		        ]
		      }
		    },
		    "notifications": {
		      "slack": {
		        "rooms": {
		          "secure": "LPNgf0Ra6Vu6I7XuK7tcnyFWJg+becx1RfAR35feWK81sru8TyuldQIt7uAKMA8tqFTP8j1Af7iz7UDokbCCfDNCX1GxdAWgXs+UKpwhO89nsidHAsCkW2lWSEM0E3xtOJDyNFoauiHxBKGKUsApJTnf39H+EW9tWrqN5W2sZg8="
		        },
		        "on_success": "never"
		      },
		      "webhooks": "https://docs.travis-ci.com/update_webhook_payload_doc"
		    },
		    "install": [
		      "rvm use 2.3.1 --install",
		      "bundle install --deployment"
		    ],
		    "script": [
		      "bundle exec rake test"
		    ],
		    ".result": "configured",
		    "global_env": [
		      "PATH=$HOME/.local/user/bin:$PATH"
		    ],
		    "group": "stable"
		  },
		  "type": "pull_request",
		  "state": "passed",
		  "status": 0,
		  "result": 0,
		  "status_message": "Passed",
		  "result_message": "Passed",
		  "started_at": "2017-10-17T20:35:49Z",
		  "finished_at": "2017-10-17T20:38:16Z",
		  "duration": 147,
		  "build_url": "https://travis-ci.org/travis-ci/docs-travis-ci-com/builds/289215159",
		  "commit_id": 84720880,
		  "commit": "729f57937e1d91106c239d3e1c12e06181df7b9e",
		  "base_commit": "729f57937e1d91106c239d3e1c12e06181df7b9e",
		  "head_commit": "735d4ff999f2e6dccc558dc179108c08fee492c2",
		  "branch": "master",
		  "message": "Merge branch 'master' into ha-git-sparse-checkout",
		  "compare_url": "https://github.com/travis-ci/docs-travis-ci-com/pull/1406",
		  "committed_at": "2017-10-17T20:34:38Z",
		  "author_name": "Hiro Asari",
		  "author_email": "asari.ruby@gmail.com",
		  "committer_name": "GitHub",
		  "committer_email": "noreply@github.com",
		  "pull_request": true,
		  "pull_request_number": 1406,
		  "pull_request_title": "Document git sparse checkout",
		  "tag": null,
		  "repository": {
		    "id": 1771959,
		    "name": "docs-travis-ci-com",
		    "owner_name": "travis-ci",
		    "url": "http://docs.travis-ci.com"
		  },
		  "matrix": [{
		    "id": 289215160,
		    "repository_id": 1771959,
		    "parent_id": 289215159,
		    "number": "2848.1",
		    "state": "passed",
		    "config": {
		      "sudo": false,
		      "dist": "trusty",
		      "language": "python",
		      "python": "3.5.2",
		      "branches": {
		        "only": [
		          "master"
		        ]
		      },
		      "cache": {
		        "pip": true,
		        "directories": [
		          "vendor/bundle",
		          "node_modules"
		        ]
		      },
		      "notifications": {
		        "slack": {
		          "rooms": {
		            "secure": "LPNgf0Ra6Vu6I7XuK7tcnyFWJg+becx1RfAR35feWK81sru8TyuldQIt7uAKMA8tqFTP8j1Af7iz7UDokbCCfDNCX1GxdAWgXs+UKpwhO89nsidHAsCkW2lWSEM0E3xtOJDyNFoauiHxBKGKUsApJTnf39H+EW9tWrqN5W2sZg8="
		          },
		          "on_success": "never"
		        },
		        "webhooks": "https://docs.travis-ci.com/update_webhook_payload_doc"
		      },
		      "install": [
		        "rvm use 2.3.1 --install",
		        "bundle install --deployment"
		      ],
		      "script": [
		        "bundle exec rake test"
		      ],
		      ".result": "configured",
		      "global_env": [
		        "PATH=$HOME/.local/user/bin:$PATH"
		      ],
		      "group": "stable",
		      "os": "linux",
		      "addons": {
		        "deploy": {
		          "provider": "heroku",
		          "api_key": {
		            "secure": "hylw2GIHMvZKOKX3uPSaLEzVrUGEA9mzGEA0s4zK37W9HJCTnvAcmgRCwOkRuC4L7R4Zshdh/CGORNnBBgh1xx5JGYwkdnqtjHuUQmWEXCusrIURu/iEBNSsZZEPK7zBuwqMHj2yRm64JfbTDJsku3xdoA5Z8XJG5AMJGKLFgUQ="
		          },
		          "app": "docs-travis-ci-com",
		          "skip_cleanup": true,
		          "true": {
		            "branch": [
		              "master"
		            ]
		          }
		        }
		      }
		    },
		    "status": 0,
		    "result": 0,
		    "commit": "729f57937e1d91106c239d3e1c12e06181df7b9e",
		    "branch": "master",
		    "message": "Merge branch 'master' into ha-git-sparse-checkout",
		    "compare_url": "https://github.com/travis-ci/docs-travis-ci-com/pull/1406",
		    "started_at": null,
		    "finished_at": null,
		    "committed_at": "2017-10-17T20:34:38Z",
		    "author_name": "Hiro Asari",
		    "author_email": "asari.ruby@gmail.com",
		    "committer_name": "GitHub",
		    "committer_email": "noreply@github.com",
		    "allow_failure": false
		  }]
		}
		`)

	var travisBody TravisWebhookBody

	err := json.Unmarshal(testString, &travisBody)

	if err == nil {
		// t.Error(travisBody)
		assertEqual(t, travisBody.ID, 289215159, "")
		assertEqual(t, travisBody.Commit, "729f57937e1d91106c239d3e1c12e06181df7b9e", "")
		assertEqual(t, travisBody.Status, 0, "")
		// log.Println(travisBody.Commit)
	} else {
		t.Fatal(err)
		// t.Fatal("Error Occured.")
	}
}
