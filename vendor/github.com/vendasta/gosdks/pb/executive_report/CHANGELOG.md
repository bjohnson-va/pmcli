## 1.10.0
- Add THERMOMETER_WITH_GRADE template type

## 1.9.1
- Fix imports to support multi package compilation

## 1.9.0
- Add endpoint for getting a list of dates with report data

## 1.8.0
- Support Fraction and time formats

## 1.7.0
- Add BASIC_NUMBER template type

## 1.6.0
- Add INFO_LIST template type

## 1.5.0
- Add BASIC_LIST template type

## 1.4.0
- Added a width field to report data
- Added link_path and link_text to report data to allow for call to action links
- Added Call To Action template type

## 1.3.0
- Raised proto number for the name changes so they don't conflict with the original parameters

## 1.2.0
- Added List endpoint
- Updated naming to proper conventions

## 1.1.0
- Added TemplateType to executive_report proto
- Added NOT_SPECIFIED as the first value in the Category enum
- Added NOT_SPECIFIED as the first value in the ReportFrequency enum
- Moved NUMBER to the first position in the Format enum
- Change value and change to be strings so we can tell if they sent 0, or if they didn't send anything

## 1.0.0
- Initial commit of executive_report protos
