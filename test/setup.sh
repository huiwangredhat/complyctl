DIR="~/.local/share/complytime/"
            CONTENT_URL="https://raw.githubusercontent.com/ComplianceAsCode/oscal-content/refs/heads/main/"
            PROFILE="fedora-cusp_fedora-default"
            CATALOG="cusp_fedora"
            COMPD="fedora-cusp_fedora-defaul"
            sudo mkdir -p  "{$DIR}/controls/"
            sudo mkdir -p  "{$DIR}/bundles/"
            sudo mkdir -p  "{$DIR}/plugins/"
            #PROFILE_URL="{$CONTENT_URL}profiles/{$PROFILE}/profile.json"
            #CATALOG_URL="{$CONTENT_URL}catalogs/{$CATALOG}/catalog.json"
            #COMPD_URL="{$CONTENT_URL}component-definitions/fedora/{$COMPD}/component-definition.json"
            sudo wget PROFILE_URL -O ${DIR}/controls/cusp_profile.json
            sudo wget CATALOG_URL -O /usr/share/complytime/controls/cusp_catalog.json
            sudo wget COMPD_URL -O /usr/share/complytime/bundles/cusp_component-definition.json
            # update trestle path
            sudo sed -i 's|trestle://catalogs/cusp_fedora/catalog.json|trestle://controls/cusp_catalog.json|' /usr/share/complytime/controls/cusp_profile.json
            sudo sed -i 's|trestle://profiles/fedora-cusp_fedora-default/profile.json|trestle://controls/cusp_profile.json|' /usr/share/complytime/bundles/cusp_component-definition.json
            sudo cp -rp bin/openscap-plugin /usr/share/complytime/plugins
            checksum=$(sha256sum /usr/share/complytime/plugins/openscap-plugin| cut -d ' ' -f 1 )
            sudo cat > /usr/share/complytime/plugins/c2p-openscap-manifest.json << EOF
            {
              "metadata": {
                "id": "openscap",
                "description": "My openscap plugin",
                "version": "0.0.1",
                "types": [
                  "pvp"
                ]
              },
              "executablePath": "openscap-plugin",
              "sha256": "$checksum",
              "configuration": [
                {
                  "name": "workspace",
                  "description": "Directory for writing plugin artifacts",
                  "required": true
                },
                {
                  "name": "profile",
                  "description": "The OpenSCAP profile to run for assessment",
                  "required": true
                },
                {
                  "name": "datastream",
                  "description": "The OpenSCAP datastream to use. If not set, the plugin will try to determine it based on system information",
                  "required": false
                },
                {
                  "name": "policy",
                  "description": "The name of the generated tailoring file",
                  "default": "tailoring_policy.xml",
                  "required": false
                },
                {
                  "name": "arf",
                  "description": "The name of the generated ARF file",
                  "default": "arf.xml",
                  "required": false
                },
                {
                  "name": "results",
                  "description": "The name of the generated results file",
                  "default": "results.xml",
                  "required": false
                }
              ]
            }
            EOF
            '
