critical:
    - rule_dsrid: DSR-1
      rule_display_id: javascript_third_parties_new_relic
      rule_description: Do not send sensitive data to New Relic.
      rule_documentation_url: https://curio.sh/reference/rules/javascript_third_parties_new_relic
      line_number: 3
      filename: pkg/commands/process/settings/rules/javascript/third_parties/new_relic/testdata/datatype_in_set_custom_attribute.js
      category_groups:
        - PII
      parent_line_number: 3
      parent_content: newrelic.setCustomAttribute("user-id", customer.email)


--
