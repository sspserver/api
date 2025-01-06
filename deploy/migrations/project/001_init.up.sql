CREATE TYPE ApproveStatus AS ENUM ('pending', 'approved', 'rejected');
CREATE TYPE ActiveStatus AS ENUM ('active', 'pause');
CREATE TYPE PrivateStatus AS ENUM ('public', 'private');
CREATE TYPE ProcessingStatus AS ENUM ('undefined', 'progress', 'processed', 'error', 'deleted');

CREATE TYPE PricingModel AS ENUM ('undefined', 'CPM', 'CPC', 'CPA');
CREATE TYPE AuctionType AS ENUM ('undefined', 'first_price', 'second_price');
CREATE TYPE FormatType AS ENUM ('undefined', 'direct', 'proxy', 'video', 'banner', 'html5', 'native', 'custom');
CREATE TYPE ZoneType AS ENUM ('zone', 'smartlink');
CREATE TYPE ApplicationType AS ENUM ('site', 'application', 'game');
CREATE TYPE PlatformType AS ENUM ('web', 'desktop', 'mobile', 'smart.phone', 'tablet', 'smart.tv', 'gamestation', 'smart.watch', 'vr', 'smart.glasses', 'smart.billboard');
CREATE TYPE RTBRequestType AS ENUM ('undefined','json','xml','protobuff','postformencoded','plain');
